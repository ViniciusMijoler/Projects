import * as React from 'react';
import MenuStore from '../main-menu/store';
import {  } from 'semantic-ui-react';
import { inject, observer } from 'mobx-react';
import NewRouterStore from '../../mobx/router.store';

interface Props {
  mainMenu?: MenuStore;
  router?: NewRouterStore;
}

@inject('mainMenu', 'router')
@observer
export default class HeaderMenu extends React.Component<Props> {

    logout = () => {
        const { setHistory } = this.props.router!;

        return setHistory('home');
    }

    render() {

        return (
            <>
            </>
        );
    }
}