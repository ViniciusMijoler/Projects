import * as React from 'react';
import { Menu, Dropdown } from 'semantic-ui-react';
import MenuStore from '../main-menu/store';
import { inject, observer } from 'mobx-react';
import NewRouterStore from '../../mobx/router.store';
import { logOff } from '../../util/auth.util';

interface Props {
  mainMenu?: MenuStore;
  router?: NewRouterStore;
}

@inject('mainMenu', 'router')
@observer
export default class HeaderMenu extends React.Component<Props> {

    logout = () => {
        const { setHistory } = this.props.router!;
        logOff();
        return setHistory('home');
    }

    render() {

        return (
            <Menu style={{ height: 50, paddingLeft: 200, marginBottom: 0 }}>
                <Menu.Menu position='right'>
                    <Dropdown text='Usuario' pointing className='link item'>
                        <Dropdown.Menu>
                            <Dropdown.Item onClick={this.logout}>Logout</Dropdown.Item>
                        </Dropdown.Menu>
                    </Dropdown>
                </Menu.Menu>
            </Menu>
        );
    }
}