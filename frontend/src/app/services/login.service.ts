import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Router} from '@angular/router';
import {} from 'jwt-decode';
import {environment} from '../../environments/environment';
import {of} from 'rxjs';
import * as moment from 'moment';


@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    private http: HttpClient,
    private router: Router
  ) { }
  getPubkey(): Observable<object>{
    return this.http.get(`/api/pubkey`);
  }

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

  public refreshSession(authResult): void {
    localStorage.setItem('idToken', authResult.token);
    localStorage.setItem('expiresAt', authResult.expiresAt);
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
    this.router.navigate(['/login']);
  }
  public isLoggedIn(): Observable<boolean>{
    return new Observable<boolean> ( observer => {
      this.refresh().subscribe((response: HttpResponse<AuthResponse>) => {
        if (response.status === 200) {
          const now = new Date();
          const expiration = new Date(response.body.expiresAt * 1000);
          if (expiration >= now) {
            observer.next(true);
          }
        }
      },
        error => {
          console.log('auth failed');
          observer.next(false);
        });
    });
  }
  refresh(): Observable<HttpResponse<any>> {
    return this.http.get(`/api/refresh`, {observe: 'response'});
  }

  isLoggedOut(): boolean {
    return !this.isLoggedIn();
  }

  hasToken(): boolean {
    const token = localStorage.getItem('idToken');
    return token !== null;
  }

  isAdmin(): Observable<boolean>{
    return new Observable<boolean> ( observer => {
      this.refresh().subscribe((response: HttpResponse<AuthResponse>) => {
        if (response.status === 200) {
          this.refreshSession(response.body);
          if (response.body.role === 'admin') {
            observer.next(true);
          }
        }
      },
        error => {
          console.log('auth failed');
          observer.next(false);
        });
    });
  }
}

class AuthResponse {
  idToken: string;
  username: string;
  email: string;
  role: string;
  expiresAt: number;
}
