import * as React from 'react';
import { withHub } from '../../Hub';

interface Props {

}
class ActionSideBar extends React.Component<Props>{
    render(){
        return (
            <div>Side Bar</div>
        )
    }
}

export default withHub(ActionSideBar);