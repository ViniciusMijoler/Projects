import { action, observable } from 'mobx';

export default class MainMenuStore {

  @observable activated: string = this.getActiveMenu();
  getActiveMenu() {
    const [, activeMenu] = window.location.pathname.split('/');
    return activeMenu || 'home';
  }

  @action setMenuActive = (name: string) => {
    this.activated = name;
  };

  routes: any[] = [
    {
      title: 'Projetos',
      description: '',
      route: 'projects',
    },
  ];

}

const mainMenu = new MainMenuStore();
export { mainMenu };