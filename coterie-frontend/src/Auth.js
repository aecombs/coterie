class Auth {
  constructor() {
    this.authenticated = false;
  }

  // login(callback) {
  //   this.authenticated = true;
  //   callback();
  // }

  // logout(callback) {
  //   this.authenticated = false;
  //   callback();
  // }

  isAuthenticated() {
    if (document.cookie.split(';').some((item) => item.trim().startsWith('session='))) {
      this.authenticated = true;
    } else {
      this.authenticated = false
    }
    
    return this.authenticated;
  }
}

export default new Auth();