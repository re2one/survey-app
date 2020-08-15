import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import {Router} from '@angular/router';
import {HttpErrorResponse} from '@angular/common/http';
import {LoginService} from '../../services/login.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  signupForm: FormGroup;

  constructor(
    public router: Router,
    private loginService: LoginService,
    private formBuilder: FormBuilder,
  ) {
    this.loginForm = this.formBuilder.group({
      email: '',
      password: ''
    });
    this.signupForm = this.formBuilder.group({
      username: '',
      email: '',
      password: '',
      passwordConfirmation: ''
    });
  }

  ngOnInit(): void {
    if (this.loginService.isLoggedIn()){
      this.router.navigate(['/surveys']);
    }
  }

  onLoginSubmit(loginData): void {
    // Process checkout data here
    // this.items = this.cartService.clearCart();
    // this.loginForm.reset();

    // console.warn('LOGIN', loginData);
    this.loginService.getAccessToken(loginData.email, loginData.password).subscribe(
      obj => {
        // this.loginData = obj;
        console.log('RESULT', obj);
        this.loginService.setSession(obj, loginData.email);
        this.router.navigate(['/surveys']);
      },
      error => console.log(error)
      // this.error.setError(error)
    );
    this.loginForm.reset();
  }

  onSignupSubmit(signupData): void {
    // Process checkout data here
    // this.items = this.cartService.clearCart();
    this.loginForm.reset();

    // console.warn('SIGNUP', signupData);
    this.loginService.signupAndGetAccessToken(signupData.username, signupData.email, signupData.password).subscribe(
      obj => {
        // this.signupData = obj;
        console.log('RESULT', obj);
        this.loginService.setSession(obj, signupData.email);
        this.router.navigate(['/surveys']);
      },
      error => console.log(error)
      // this.error.setError(error)
    );
  }
}
