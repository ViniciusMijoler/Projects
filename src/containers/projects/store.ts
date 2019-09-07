import { observable } from 'mobx';

export default class ProjectsStore {
  @observable records: any[] = [];
}
const projects = new ProjectsStore();
export { projects };
