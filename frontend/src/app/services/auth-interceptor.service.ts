import { Injectable } from '@angular/core';import {HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from "@angular/common/http";
import {Observable} from 'rxjs';
import {LoginService} from './login.service';
import {tap} from 'rxjs/operators';
import {Router} from '@angular/router';
// import {OnofflineService} from './onoffline.service';
import {environment} from '../../environments/environment';


@Injectable({
  providedIn: 'root'
})

export class AuthInterceptor {

  // constructor(private loginService: LoginService, private router: Router, private onofflineService: OnofflineService) {
  constructor(private loginService: LoginService, private router: Router) {

  }

  intercept(req: HttpRequest<any>,
            next: HttpHandler): Observable<HttpEvent<any>> {

    const idToken = localStorage.getItem('idToken');

    if (idToken) {
      const cloned = req.clone({
        headers: req.headers.set('Authorization',
          'Bearer ' + idToken)
      });
      /*
      if (environment.debug)
        console.log('HTTP INTERCEPTOR REQ > ', req);
        */
      return next.handle(cloned).pipe(tap((error: HttpEvent<any>) => {
        /*
        if (environment.debug)
          console.log('HTTP INTERCEPTOR RES < ', error);
         */
        if (error instanceof HttpErrorResponse) {
          if (error.status === 401) {
            this.router.navigate(['/login']);
          }
        }
      }));
    } else {
      return next.handle(req);
    }
  }

}
