import Header from "./components/Header";

class Auth {
  constructor() {
    this.authenticated = false;
  }

  login(callback) {
    this.authenticated = true;
    Header.setLoginButtonText("Logout");
    callback();
  }

  logout(callback) {
    this.authenticated = false;
    Header.setLoginButtonText("Login");
    callback();
  }

  isAuthenticated() {
    return this.authenticated;
  }
}

export default new Auth();