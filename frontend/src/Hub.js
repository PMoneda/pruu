/* eslint-disable react/display-name */
import React, { Component } from "react";

//Reference: https://en.wikipedia.org/wiki/Observer_pattern;

export class Hub {
  constructor() {
    this.observers = {};
    this.buffer = {};
    this.mailboxes = {};
    this.isReplaying = false;
  }
  /**
   *
   * @param {string} event Nome do evento que será escutado
   * @param {function} observer Função de callback que será chamada
   * @param {boolean} withLastValue Caso seja true durante o attach caso tenha alguma mensagem na fila
   *  a última já é passada para o observer
   */
  attach(event, observer, withLastValue) {
    if (!this.observers[event]) {
      this.observers[event] = [];
      this.buffer[event] = [];
    }
    if (typeof observer === "undefined") {
      return;
    }
    this.observers[event].push(observer);
    if (withLastValue === true) {
      //Caso já tenha algo na fila já dispara o callback
      if (this.buffer[event] && this.buffer[event].length > 0) {
        setTimeout(() => {
          const last = this.buffer[event].length - 1;
          observer(this.buffer[event][last]);
        }, 1);
      }
    }
  }

  detach(event, observer) {
    if (!this.observers[event]) {
      return;
    }
    let index = -1;
    for (let i = 0; i < this.observers[event].length; i++) {
      if (this.observers[event][i] === observer) {
        index = i;
        break;
      }
    }
    if (index >= 0) {
      this.observers[event].splice(index, 1);
    }
  }
  detachAll(event) {
    if (!this.observers[event]) {
      return;
    }
    this.observers[event].length = 0;
  }
  detachAllWithPrefix(prefix) {
    const events = Object.keys(this.observers).filter(s =>
      s.startsWith(prefix)
    );
    events.forEach(e => (this.observers[e].length = 0));
  }

  attachMailbox(address, callback) {
    if (!this.mailboxes[address]) {
      this.mailboxes[address] = [];
    }
    this.mailboxes[address].push(callback);
  }

  request(address, message) {
    if (!this.mailboxes[address]) {
      return new Promise((res, rej) => rej("address not registered!"));
    }
    return Promise.all(this.mailboxes[address].map(x => x(message)));
  }

  notify(event, message) {
    if (!this.observers[event]) {
      return;
    }
    this.buffer[event].pop();
    this.buffer[event].push(message);
    this.observers[event].forEach(observer => {
      setTimeout(function() {
        if (typeof observer === "function") {
          observer(message);
        }
      }, 1);
    });
  }
  replay(event, rollback) {
    if (this.buffer[event]) {
      const last = this.buffer.length - 1;
      for (let i = last - rollback; i <= last; i++) {
        this.notify(event, this.buffer[event][i]);
      }
    }
  }
}

export function withHub(WrappedComponent) {
  return class extends Component {
    constructor(props) {
      super(props);
      this.attach = this.attach.bind(this);
      this.notify = this.notify.bind(this);
      this.toDetach = {};
    }

    attach(event, subscriber, withInitialValue) {
      if (!this.toDetach[event]) {
        this.toDetach[event] = [];
      }
      this.toDetach[event].push(subscriber);
      window.Hub.attach(event, subscriber, withInitialValue);
    }

    detach(event, subscriber) {
      window.Hub.detach(event, subscriber);
    }

    notify(event, message) {
      window.Hub.notify(event, message);
    }

    request(address, message) {
      return window.Hub.request(address, message);
    }

    componentWillUnmount() {
      Object.keys(this.toDetach).map(event => {
        this.toDetach[event].forEach(e => {
          window.Hub.detach(event, e);
        });
      });
    }

    render() {
      return (
        <WrappedComponent
          {...this.props}
          attach={this.attach}
          notify={this.notify}
          detach={this.detach}
          request={this.request}
        />
      );
    }
  };
}