import { observable, action } from 'mobx';
import Swal from 'sweetalert2';
import { setAuth } from '../../util/auth.util';
import { assign } from '../../util';
import { success, warning } from '../../components/notifications';

export default class LoginStore {
  @observable user_name = '';
  @observable password = '';
  @observable loading = false;

  @action handleChange = (event: any, select?: any) => {
    const { id, value } = select || event.target;
    assign(this, id, value);
  }

  @action handleSubmit = async () => {
    const { user_name, password } = this;
    this.loading = true
    try {
      if (user_name === 'tatiana@gmail.com' && password === '123') {
        setAuth("12345678");
        success("Seja Bem Vindo!")
      } else {
        warning("Usuário ou Senha incorreta, tente novamente!");
      }
      this.loading = false;
    } catch (error) {
      this.loading = false;
      Swal.fire({
        text: 'Ocorreu um erro não esperado.',
        type: 'error'
      });
      throw error;
    }
  }
}
const login = new LoginStore();
export { login };