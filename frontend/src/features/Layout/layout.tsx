import * as React from "react"
import AppHeader from "../AppHeader/appHeader";
import ActionSideBar from "../ActionSideBar/actionSideBar";
import Explorer from "../Explorer/explorer";
import TabPanel from "../TabPanel/tabPanel";
import { withHub } from "../../Hub";


class Layout extends React.Component {
    render(){
        const root:React.CSSProperties = {
            display:"flex",
        }
        const header:React.CSSProperties = {
            flex:1,
        }
        const appheader:React.CSSProperties = {
            width:"100%",
            height:"40px",
            backgroundColor:"#335792",
            color:"white"
        }
        const appbottom:React.CSSProperties = {
            width:"100%",
            height:"30px",
            backgroundColor:"#212121",
        }
        const actionsidebar:React.CSSProperties = {
            width:"60px",
            height:`${window.innerHeight-70}px`,
            backgroundColor:"#212121",
            textAlign:"center",
            boxSizing:"border-box",
            flexBasis:"auto"
        }
        const explorer:React.CSSProperties = {
            flex:1,
            maxWidth:"300px",
            height:`${window.innerHeight-70}px`,
            backgroundColor:"#28292b",
            textAlign:"center",
            boxSizing:"border-box",
            flexBasis:"auto"
        }
        
        const tabpanel:React.CSSProperties = {
            flex:4,
            backgroundColor:"#0e0e0e"
        }
        return (
            <div style={root} >
                <div style={header}>
                   <div style={appheader}>Pruu Log4dev</div> 
                   <div style={root}>
                        <div style={actionsidebar}><ActionSideBar/></div>
                        <div style={explorer}><Explorer/></div>
                        <div style={tabpanel}><TabPanel/></div>
                   </div>                   
                   <div style={appbottom}>D</div> 
                </div>                
            </div>
        )
    }
}

export default withHub(Layout)