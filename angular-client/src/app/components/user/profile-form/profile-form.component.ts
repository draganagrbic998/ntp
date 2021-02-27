import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SNACKBAR_CLOSE, SNACKBAR_ERROR, SNACKBAR_ERROR_OPTIONS, SNACKBAR_SUCCESS_OPTIONS } from 'src/app/constants/snackbar';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-profile-form',
  templateUrl: './profile-form.component.html',
  styleUrls: ['./profile-form.component.scss']
})
export class ProfileFormComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private userService: UserService,
    private snackBar: MatSnackBar,
    public location: Location
  ) { }

  savePending = false;
  profileForm: FormGroup = new FormGroup({
    email: new FormControl(this.authService.getUser().email, [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    first_name: new FormControl(this.authService.getUser().first_name, [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    last_name: new FormControl(this.authService.getUser().last_name, [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    address: new FormControl(this.authService.getUser().address, [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    city: new FormControl(this.authService.getUser().city, [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    zip_code: new FormControl(this.authService.getUser().zip_code, [Validators.required, Validators.pattern(new RegExp('\\S'))]),

    old_password: new FormControl(''),
    new_password: new FormControl(''),
    new_password_confirmation: new FormControl('')
  }, {
    validators: [this.passwordConfirmed()]
  });

  resetPassword(): void{
    this.profileForm.get('old_password').setValue('');
    this.profileForm.get('new_password').setValue('');
    this.profileForm.get('new_password_confirmation').setValue('');
  }

  save(): void{
    if (this.profileForm.invalid){
      return;
    }
    this.savePending = true;
    this.userService.save(this.profileForm.value).subscribe(
      (user: User) => {
        this.savePending = false;
        this.resetPassword();
        if (user){
          this.authService.saveUser(user);
          this.snackBar.open('Profile updated!', SNACKBAR_CLOSE, SNACKBAR_SUCCESS_OPTIONS);
        }
        else{
          this.snackBar.open(SNACKBAR_ERROR, SNACKBAR_CLOSE, SNACKBAR_ERROR_OPTIONS);
        }
      }
    );
  }

  ngOnInit(): void {
  }

  private passwordConfirmed(): ValidatorFn{
    return (control: AbstractControl): ValidationErrors => {
      if (!control.get('old_password').value && (control.get('new_password').value || control.get('new_password_confirmation').value)){
        return {oldPasswordError: true};
      }
      if (control.get('new_password').value !== control.get('new_password_confirmation').value){
        return {newPasswordError: true};
      }
      return null;
    };
  }

}
