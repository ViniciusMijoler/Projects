import * as React from 'react';
import { Container, Header, Segment, Form, FormGroup, Icon, Button } from 'semantic-ui-react';
import { inject, observer } from 'mobx-react';
import ReactDatePicker from "react-datepicker";
import ProjectsStore from './store';
import './index.css';

interface Props {
  projects: ProjectsStore
}

@inject('projects')
@observer
export default class Projects extends React.Component<Props> {
  
  render() {
    return (
      <Container  style={{ paddingTop: 15, paddingBottom: 15 }}>
          <Header as='h2'>
              Projetos
          </Header>
          <Segment>
            <Form>
              <FormGroup widths='equal'>
                <Form.Field>
                  <Form.Input
                    id="nome_projeto" 
                    label='Nome projeto'
                    value={''}
                    />
                </Form.Field>

                <Form.Field>
                  <Form.Input
                    id="nome_empresa" 
                    label='Nome empresa'
                    value={''}
                    />
                </Form.Field>
                
                <Form.Field>
                  <Form.Input
                    id="palavras_chave" 
                    label='Palavras chave'
                    value={''}
                    />
                </Form.Field>
              </FormGroup>
              
              <FormGroup widths='3'>
                <Form.Field>
                  <Form.Input
                    id="area_proejto" 
                    label='Area projeto'
                    value={''}
                    />
                </Form.Field>

                <Form.Field>
                  <label>Data Limite</label>
                  <ReactDatePicker
                    value={(new Date().toISOString()).split('T')[0].split('-').reverse().join('/')}
                    isClearable
                    dateFormat='dd/MM/yyyy'
                    showYearDropdown
                    showMonthDropdown/>
                </Form.Field>
                
              </FormGroup>

              <Form.Group className='row-reverse'>
                <Form.Field className='no-label'>
                  <Button
                    title="Pesquisar"
                    type='submit'
                    floated='right'
                    icon
                    labelPosition='left'
                    color='green'
                    size='small'>
                    <Icon name='save' />
                    Pesquisar
                  </Button>
                </Form.Field>
              </Form.Group>
            </Form>
          </Segment>
      </Container>
    );

  }
}
