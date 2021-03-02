import { Injectable } from '@angular/core';
import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

@Injectable({
  providedIn: 'root'
})
export class DateValidatorService {

  constructor() { }

  dateValidator(): ValidatorFn{
    return (control: AbstractControl): ValidationErrors => {
      let dateValid = true;
      if (control.value < new Date()){
        dateValid = false;
      }
      return dateValid ? null : {dateError: true};
    };
  }

}
