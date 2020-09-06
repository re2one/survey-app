import { Injectable } from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from '@angular/router';
import {LoginService} from './login.service';
import {Observable, of} from 'rxjs';
import {catchError, map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class RoleGuardService implements CanActivate {

  constructor(private loginService: LoginService, private router: Router) { }
  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean>{
    return this.loginService.isAdmin().pipe(
      map(
        response => {
          if (response) {
            return true;
          } else {
            this.router.navigate(['/surveys']);
            return false;
          }
        }),
      catchError((err, response) => {
        this.router.navigate(['/surveys']);
        return of(false);
      })
    );
  }
}
