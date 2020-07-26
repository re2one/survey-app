import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  signupForm: FormGroup;

  constructor(
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
  }

  onLoginSubmit(loginData): void {
    // Process checkout data here
    // this.items = this.cartService.clearCart();
    this.loginForm.reset();

    console.warn('LOGIN', loginData);
  }

  onSignupSubmit(signupData): void {
    // Process checkout data here
    // this.items = this.cartService.clearCart();
    this.loginForm.reset();

    console.warn('SIGNUP', signupData);
  }
}
