import * as React from 'react';
import { Form, FormGroup, Button, Header, Segment } from 'semantic-ui-react';
import { observer } from 'mobx-react';
import ProjectsStore from './store';
import ReactDatePicker from "react-datepicker";
import { getDate } from '../../util';

interface Props {
  projects: ProjectsStore
}

@observer
export default class NewProject extends React.Component<Props>{

    handleSubmit = async (e: any) => {
        e.preventDefault();
        const { handleSubmit } = this.props.projects;
        handleSubmit();
    }

    render() {
        const {
            toggleScreen,
            handleChangeNew,
            handleDateNew,
            newProject,
            isLoading
        } = this.props.projects
        return (
            <>
                <Header as="h2">
                    Projetos
                </Header>
                <Segment>
                    <Header as='h2'>
                        Novo
                    </Header>
                    <Form onSubmit={this.handleSubmit}>
                        <FormGroup widths='equal'>
                            <Form.Field>
                                <Form.Input
                                    fluid
                                    id="nome"
                                    label='Nome projeto'
                                    value={newProject.nome}
                                    onChange={handleChangeNew}/>
                            </Form.Field>

                            <Form.Field>
                                <label>Data Limite</label>
                                <ReactDatePicker
                                    id="data_limite"
                                    selected={getDate(newProject.data_limite)}
                                    isClearable
                                    value={newProject.data_limite}
                                    dateFormat='dd/MM/yyyy'
                                    onChange={(date: any) => handleDateNew(date, 'data_limite')}
                                    showYearDropdown
                                    showMonthDropdown/>
                            </Form.Field>

                            <Form.Field>
                                <Form.Input
                                    fluid
                                    id="palavras_chaves" 
                                    label='Palavras chave'
                                    value={newProject.palavras_chaves}
                                    onChange={handleChangeNew}/>
                            </Form.Field>
                            
                            <Form.Field>
                                <Form.Input
                                    id="area_projeto" 
                                    label='Area projeto'
                                    value={newProject.area_projeto}
                                    onChange={handleChangeNew}/>
                            </Form.Field>
                        </FormGroup>

                        <FormGroup widths="equal">
                            <Form.Field>
                                <Form.TextArea
                                    id="descricao" 
                                    label='Descriçao'
                                    value={newProject.descricao}
                                    onChange={handleChangeNew}/>
                            </Form.Field>
                        </FormGroup>

                        <Form.Group style={{ flexDirection: 'row-reverse' }}>
                            <Form.Field>
                                <Button.Group>
                                    <Button onClick={toggleScreen}>Cancelar</Button>
                                    <Button.Or text='ou' />
                                    <Button positive type='submit' loading={isLoading} onClick={this.handleSubmit}>Salvar</Button>
                                </Button.Group>
                            </Form.Field>
                        </Form.Group>
                    </Form>
                </Segment>
            </>
        )
    }
}