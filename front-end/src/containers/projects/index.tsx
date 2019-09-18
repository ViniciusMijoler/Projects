import * as React from 'react';
import { Container, Header, Grid } from 'semantic-ui-react';
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
      <Container>
          <Header as='h2'>Projetos</Header>
          <Grid.Column>
          </Grid.Column>
      </Container>
    );

  }
}
