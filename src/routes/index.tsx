import * as React from 'react';
import { observer } from 'mobx-react';
import {
  Route,
  Switch,
  withRouter,
} from 'react-router-dom';
import NotFound from '../containers/not-found';
import MainMenu from '../components/main-menu';
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
          <div style={{ minHeight: '100%'}}>
            <MainMenu />
            <div style={{ paddingTop: 46, minHeight: '100vh' }}>
              <PrivateRoutes />
            </div>
          </div>
        </>
        <Route render={props => <NotFound {...props} />} />
      </Switch>
    );
  }
}