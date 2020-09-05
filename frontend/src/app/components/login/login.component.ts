import { Component, OnInit } from '@angular/core';
import {FormGroup, FormBuilder, ValidationErrors, ValidatorFn} from '@angular/forms';
import {Router} from '@angular/router';
import {Validators} from '@angular/forms';
import {MatError} from '@angular/material/form-field';
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
      email: ['', [Validators.required]],
      password: ['', [Validators.required]],
    });
    this.signupForm = this.formBuilder.group({
      username: ['', [Validators.required]],
      email: ['', [Validators.required]],
      password: ['', [Validators.required]],
      passwordConfirmation: ['', [Validators.required]],
    }, {validators: this.checkPasswords });
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

  checkPasswords: ValidatorFn = (control: FormGroup): ValidationErrors | null => {// here we have the 'passwords' group
    const pass = control.get('password').value;
    const confirmPass = control.get('passwordConfirmation').value;

    return pass === confirmPass ? null : {passwordsEqual: true};
  }
}
