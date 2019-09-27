import * as React from 'react';
// import { Container, Card, Header } from 'semantic-ui-react';
import { inject, observer } from 'mobx-react';
import LoginStore from './store';

interface Props {
  login: LoginStore;
}

@inject('login')
@observer
export default class Login extends React.Component<Props> {

  render() {
    return (
        <section className={'login-register login-sidebar'}>
          <div className={'login-box'}>

          </div>
        </section>
    );

  }
}
