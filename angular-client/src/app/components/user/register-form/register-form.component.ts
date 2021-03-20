import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SNACKBAR_CLOSE, SNACKBAR_ERROR, SNACKBAR_ERROR_OPTIONS, SNACKBAR_SUCCESS_OPTIONS } from 'src/app/constants/snackbar';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-register-form',
  templateUrl: './register-form.component.html',
  styleUrls: ['./register-form.component.scss']
})
export class RegisterFormComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private userService: UserService,
    private snackBar: MatSnackBar
  ) { }

  registerPending = false;
  registerForm: FormGroup = new FormGroup({
    first_name: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    last_name: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    address: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    city: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    zip_code: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    email: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    password: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    password_confirmation: new FormControl('', [this.passwordConfirmed()])
  });

  register(): void{
    if (this.registerForm.invalid){
      return;
    }
    this.registerPending = true;
    this.authService.deleteUser();
    // tslint:disable-next-line: deprecation
    this.userService.register(this.registerForm.value).subscribe(
      (response: boolean) => {
        this.registerPending = false;
        if (response){
          this.registerForm.reset();
          this.snackBar.open('We have sent you an email with verification in it!',
          SNACKBAR_CLOSE, SNACKBAR_SUCCESS_OPTIONS);
        }
        else{
          this.snackBar.open(SNACKBAR_ERROR, SNACKBAR_CLOSE, SNACKBAR_ERROR_OPTIONS);
        }
      }
    );
  }

  ngOnInit(): void {
    // tslint:disable-next-line: deprecation
    this.registerForm.get('password').valueChanges.subscribe(
      () => {
        this.registerForm.get('password_confirmation').updateValueAndValidity();
      }
    );
  }

  private passwordConfirmed(): ValidatorFn{
    return (control: AbstractControl): ValidationErrors => {
      const passwordConfirmed: boolean = control.parent ?
      control.value === control.parent.get('password').value : true;
      return passwordConfirmed ? null : {passwordError: true};
    };
  }

}
