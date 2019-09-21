import * as React from 'react';
import { Header } from 'semantic-ui-react';
import { inject, observer } from 'mobx-react';
import ProjectsStore from './store';

interface Props {
    projects: ProjectsStore
}

@inject('projects')
@observer
export default class Projects extends React.Component<Props> {
  
  render() {
    return (
      <div style={{ paddingTop: 30, paddingLeft: 40, paddingRight: 40, paddingBottom: 30 }}>
          <Header as='h2'>Projetos</Header>
          <div style={{ backgroundColor: '#fff', padding: 20 }} >
            <Header as='h3'>Filtros</Header>
            
          </div>
      </div>
    );

  }
}
