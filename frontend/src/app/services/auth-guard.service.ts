import { Injectable } from '@angular/core';
import { CanActivate, CanActivateChild, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import {LoginService} from './login.service';
import {Observable, of} from 'rxjs';
import {map, catchError} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate{

  constructor(private loginService: LoginService, private router: Router) { }
  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean>{
    return this.loginService.isLoggedIn().pipe(
      map(
        response => {
          if (response) {
            return true;
          } else {
            return false;
          }
        })
      ,
      catchError((err, response) => {
        this.router.navigate(['/login']);
        return of(false);
      })
    );
  }
}
