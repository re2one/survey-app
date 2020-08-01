import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {of} from 'rxjs';
import * as moment from 'moment';


@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    private http: HttpClient
  ) { }

  getAccessToken(email: String, password: String) {
      return this.http.post(`/api/login`, {
        'email': email,
        'password': password,
    });
  }

  signupAndGetAccessToken(username:String, email: String, password: String) {
      return this.http.post(`/api/signup`, {
        'email': email,
        'password': password,
        'name': username
      });
  }

  public setSession(authResult, email: string) {
    localStorage.setItem('idToken', authResult.idToken);
    localStorage.setItem('expiresAt', authResult.expiresAt);
    localStorage.setItem('username', email);
    localStorage.setItem('admin', authResult.admin);
    // this.username = email;
  }

  logout() {
    localStorage.removeItem('idToken');
    localStorage.removeItem('expiresAt');
    localStorage.removeItem('username');
    localStorage.removeItem('admin');
    // this.username = null;
  }

  // getUserName() {
  //   if (this.username == null) {
  //     return (this.username = localStorage.getItem('username'));
  //   }
  //   return this.username;
  // }

  public isLoggedIn() {
    // return localStorage.getItem('id_token') !== null;
    return moment().isBefore(this.getExpiration());
  }

  isLoggedOut() {
    return !this.isLoggedIn();
  }

  getExpiration() {
    const expiration = localStorage.getItem('expiresAt');
    return moment(expiration);
  }

  isAdmin(): boolean {
    return (localStorage.getItem('admin') || '').toLowerCase() === 'true';
  }
}
