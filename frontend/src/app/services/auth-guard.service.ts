import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot} from '@angular/router';
import {LoginService} from './login.service';
import {Observable, of} from 'rxjs';
import {map, catchError} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate {

  constructor(private loginService: LoginService) { }
  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean>{
    return this.loginService.isLoggedIn().pipe(
      map(
        response => {
          if (response) {
            return true;
          } else {
            this.loginService.logout();
            return false;
          }
        })
      ,
      catchError((err, response) => {
        this.loginService.logout();
        return of(false);
      })
    );
  }
}
