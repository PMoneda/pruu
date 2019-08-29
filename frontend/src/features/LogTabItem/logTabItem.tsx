import React from "react";
import TabItem from "../../core/tabItem"
import { withHub } from "../../Hub";

export interface IProps {
    editable:boolean
}

export interface IState {
    editable:boolean
    variable:string
}

class LogTabItem extends React.Component<IProps,IState> implements TabItem {
    
    constructor(props:IProps){
        super(props);
        this.state = {
            editable:props.editable,
            variable:""
        }
        this.onEnter = this.onEnter.bind(this)
    }
    
    onEnter(evt:React.KeyboardEvent){
        
        const t = evt.target as any;
        if(evt.keyCode===13){
            //emitir o t.value

            
        }
        
    }

    header(): React.ReactNode {
        if(this.state.editable){
            return <div><input onKeyUp={this.onEnter} type="text" placeholder="value"/></div>
        }
        return <div>Titulo</div>
    }
    
    render(){
        return this.header();
    }
    
    body(): React.ReactNode {
        return <div>
            <div>Log1</div>
            <div>Log2</div>
            <div>Log3</div>
        </div>
    }

}

export default  LogTabItem