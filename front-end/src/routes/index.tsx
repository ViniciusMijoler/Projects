import * as React from 'react';
import { observer } from 'mobx-react';
import {
  Route,
  Switch,
  withRouter,
} from 'react-router-dom';
import NotFound from '../containers/not-found';
import MainMenu from '../components/main-menu';
import HeaderMenu from '../components/header-menu';
import PrivateRoutes from './private-routes';

// @ts-ignore
@withRouter
@observer
export default class Routes extends React.Component {

  render() {
    // const publicUrl = process.env.PUBLIC_URL;

    return (
      <Switch>
        <>
        <div style={{ paddingTop: 50 }}>
          <HeaderMenu />
          <MainMenu />
          <div style={{ paddingLeft: 200, position: 'fixed', height: '100%', width: '100%', background: '#f0f0f7', overflowX: 'auto' }}>
            <div style={{ position: 'relative' }}>
              <PrivateRoutes />
            </div>
          </div>
        </div>
        </>
        <Route render={props => <NotFound {...props} />} />
      </Switch>
    );
  }
}