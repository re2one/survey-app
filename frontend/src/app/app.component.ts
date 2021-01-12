import {Component} from '@angular/core';
import {LoginService} from './services/login.service';
import {Router} from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  constructor(
    private loginService: LoginService,
    private router: Router
  ) {
  }

  isLoggedIn(): boolean {
    return this.loginService.hasToken();
  }

  logout(): void {
    this.loginService.logout();
  }

  about(): void {
    this.router.navigate(['about']);
  }
}
