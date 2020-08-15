import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
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

  getAccessToken(email: string, password: string): Observable<object> {
      return this.http.post(`/api/login`, {
        email,
        password,
    });
  }

  signupAndGetAccessToken(username: string, email: string, password: string): Observable<object> {
      return this.http.post(`/api/signup`, {
        email,
        password,
        name: username
      });
  }

  public setSession(authResult, email: string): void {
    localStorage.setItem('idToken', authResult.token);
    localStorage.setItem('expiresAt', authResult.expiresAt);
    localStorage.setItem('email', email);
    localStorage.setItem('username', authResult.username);
    localStorage.setItem('role', authResult.role);
    // this.username = email;
  }

  logout(): void {
    localStorage.removeItem('idToken');
    localStorage.removeItem('expiresAt');
    localStorage.removeItem('username');
    localStorage.removeItem('role');
    localStorage.removeItem('email');
    // this.username = null;
  }

  // getUserName() {
  //   if (this.username == null) {
  //     return (this.username = localStorage.getItem('username'));
  //   }
  //   return this.username;
  // }

  public isLoggedIn(): boolean {
    // return localStorage.getItem('id_token') !== null;
    // return moment().isBefore(this.getExpiration());
    const now = new Date();
    const expiration = new Date(this.getExpiration() * 1000);
    return now < expiration;
  }

  isLoggedOut(): boolean {
    return !this.isLoggedIn();
  }
  getExpiration(): number {
  // getExpiration(): moment.Moment {
    const expiration = localStorage.getItem('expiresAt');
    // return moment(expiration);
    return parseInt(expiration, 10);
  }

  isAdmin(): boolean {
    return (localStorage.getItem('admin') || '').toLowerCase() === 'true';
  }
}
