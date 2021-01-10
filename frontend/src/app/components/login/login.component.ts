import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, ValidationErrors, ValidatorFn, Validators} from '@angular/forms';
import {Router} from '@angular/router';
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
      wantsThesis: ['', [Validators.required]],
    }, {validators: this.checkPasswords });
  }

  ngOnInit(): void {
    if (this.loginService.isLoggedIn()){
      this.router.navigate(['/surveys']);
    }
  }

  onLoginSubmit(loginData): void {

    this.loginService.getAccessToken(loginData.email, loginData.password).subscribe(
      obj => {
        this.loginService.setSession(obj, loginData.email);
        this.router.navigate(['/surveys']);
      },
      error => console.log(error)
    );
    this.loginForm.reset();
  }

  onSignupSubmit(signupData): void {
    this.loginForm.reset();

    this.loginService.signupAndGetAccessToken(signupData.username, signupData.email, signupData.wantsThesis, signupData.password).subscribe(
      obj => {
        this.loginService.setSession(obj, signupData.email);
        this.router.navigate(['/surveys']);
      },
      error => console.log(error)
    );
  }

  checkPasswords: ValidatorFn = (control: FormGroup): ValidationErrors | null => {// here we have the 'passwords' group
    const pass = control.get('password').value;
    const confirmPass = control.get('passwordConfirmation').value;

    return pass === confirmPass ? null : {passwordsEqual: true};
  }
}
