import * as React from "react"
import { withHub } from "../../Hub";
import TabList from "../../core/tabList";
import LogTabItem from "../LogTabItem/logTabItem";

const tabheader:React.CSSProperties = {
    display:"flex",    
    backgroundColor:"#212121",
    textAlign:"center",
    boxSizing:"border-box",
    flexBasis:"auto",
    
    
}

const tabheaderitem:React.CSSProperties = {
    
    backgroundColor:"#212121",
    textAlign:"center",
    boxSizing:"border-box",
    flexBasis:"auto",
    color:"gray",
    minWidth:"100px",
    paddingLeft:"10px",
    paddingRight:"10px",
    paddingBottom:"5px",
    paddingTop:"5px",
    display:"flex",
    borderRight:"2px solid black"
}
const tabheaderadditem:React.CSSProperties = {
    
    backgroundColor:"#161616",
    textAlign:"center",
    boxSizing:"border-box",
    flexBasis:"auto",
    color:"gray",
    minWidth:"30px",
    paddingLeft:"10px",
    paddingRight:"10px",
    paddingBottom:"5px",
    paddingTop:"5px",
    borderRight:"2px solid black",
    cursor:"pointer"
}

const headerclosebutton:React.CSSProperties = {
    textAlign:"center",
    width:"20px",
    cursor:"pointer",
}
export interface IProps {    
}
export interface IState {
    tabs: TabList;
}
class TabPanel extends React.Component<IProps,IState> {

    constructor(props:IProps){
        super(props);
        this.state = {
            tabs:new TabList()
        }
        this.addNewTab = this.addNewTab.bind(this);
        this.closeTabItem = this.closeTabItem.bind(this);
    }
    addNewTab(){
        const tabItem = new LogTabItem({editable:true});
        this.setState(x => {
            x.tabs.addTabItem(tabItem);
            x.tabs.setSelectedTab(x.tabs.getLastTabIndex())
            return x
        });
    }
    closeTabItem(index:number){
        this.setState(x => {
            console.log("removing tab "+index)
            x.tabs.removeTabItem(index);
            return x;
        })
    }

    renderHeaderItem(x:React.ReactNode,index:number):React.ReactNode {
        return <div key={index} style={tabheaderitem}>
                 <div>{x}</div>
                 <div style={headerclosebutton} onClick={()=>this.closeTabItem(index)}>x</div>
               </div>
    }

    render(){
        return (
            <div>
                <div style={tabheader}>
                    {this.state.tabs.getHeaders().map((x,i)=>this.renderHeaderItem(x,i))}
                    <div style={tabheaderadditem} onClick={this.addNewTab}>+</div>
                </div>
                <div>Hello Moneda</div>
            </div>
        )
    }
}

export default withHub(TabPanel);