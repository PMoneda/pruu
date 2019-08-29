import * as React from "react"
import { withHub } from "../../Hub";

class AppHeader extends React.Component {
    render(){
        return (
            <div>Header</div>
        )
    }
}


export default withHub(AppHeader);