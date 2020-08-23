import { Component, OnChanges } from '@angular/core';
import {MatToolbarModule} from '@angular/material/toolbar';
import {AuthTestService} from './services/auth-test.service';
import {LoginService} from './services/login.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  constructor(
    private loginService: LoginService
  ) { }

  isLoggedIn(): boolean{
    return this.loginService.hasToken();
  }
  logout(): void{
    this.loginService.logout();
  }
}
