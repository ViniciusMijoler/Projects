import { action, observable } from 'mobx';
import { NewProjectInterface } from '../../interfaces/project.interface';
import { assign } from '../../util';
import { postProject } from '../../api/projects.api';

const initialNewProject = {
  nome: '',
  id_empresa: 1,
  palavras_chaves: '',
  area_projeto: '',
  data_limite: '',
  descricao: ''
}

const initialFilter = {
  nome_projeto: '',
  nome_empresa: '',
  palavras_chave: '',
  area_projeto: '',
  data: '',
}

export default class ProjectsStore {

  @observable newProject: NewProjectInterface = initialNewProject;

  @observable filter: {
    nome_projeto: string;
    nome_empresa: string;
    palavras_chave: string;
    area_projeto: string;
    data: string;
  } = initialFilter;

  @observable records: any[] = [];

  @observable showNewProjectScreen: boolean = true;

  @action toggleScreen = () => {
    this.showNewProjectScreen = !this.showNewProjectScreen;
    if (this.showNewProjectScreen) {
      this.filter = initialFilter;
    } else {
      this.newProject = initialNewProject;
    }
  }

  @action handleChangeNew = (event: any, select?: any) => {
    const { id, value } = select || event.target;
    assign(this.newProject, id, value);
  };

  @action handleChangeFilter = (event: any, select?: any) => {
    const { id, value } = select || event.target;
    assign(this.filter, id, value);
  };

  @action handleDateNew = (data: Date|null, id: string) => {
    assign(this.newProject, id, data ? data.toISOString().split('T')[0].split('-').reverse().join('/') : '');
  }

  @action handleDateFilter = (data: Date|null, id: string) => {
    assign(this.filter, id, data ? data.toISOString().split('T')[0].split('-').reverse().join('/') : '');
  }

  @action handleSubmit = () => {
    const data = { ...this.newProject }
    postProject(data)
      .then((res) => {
        console.log(res);
        alert('salvo');
      })
      .catch(err => {
        console.error(err);
      })
  }

}
const projects = new ProjectsStore();
export { projects };
