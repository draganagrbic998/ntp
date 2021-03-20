import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SNACKBAR_CLOSE, SNACKBAR_ERROR, SNACKBAR_ERROR_OPTIONS, SNACKBAR_SUCCESS_OPTIONS } from 'src/app/constants/snackbar';
import { Advertisement } from 'src/app/models/ad';
import { Image } from 'src/app/models/image';
import { AdService } from 'src/app/services/ad/ad.service';

@Component({
  selector: 'app-ad-form',
  templateUrl: './ad-form.component.html',
  styleUrls: ['./ad-form.component.scss']
})
export class AdFormComponent implements OnInit {

  constructor(
    private adService: AdService,
    private snackBar: MatSnackBar,
    public location: Location
  ) { }

  savePending = false;
  adForm: FormGroup = new FormGroup({
    Category: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    Name: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    Description: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    Price: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S')), Validators.pattern(/^[0-9]\d*$/)]),
  });
  images: Image[] = [];

  save(): void{
    if (this.adForm.invalid){
      return;
    }
    this.savePending = true;
    // tslint:disable-next-line: deprecation
    this.adService.save({...this.adForm.value, ...{Images: this.images, ID: this.adService.selectedAd?.ID}}).subscribe(
      (ad: Advertisement) => {
        this.savePending = false;
        if (ad){
          this.snackBar.open('Advertisement saved!', SNACKBAR_CLOSE, SNACKBAR_SUCCESS_OPTIONS);
          this.adService.announceRefreshData();
        }
        else{
          this.snackBar.open(SNACKBAR_ERROR, SNACKBAR_CLOSE, SNACKBAR_ERROR_OPTIONS);
        }
      }
    );
  }

  get edit(): boolean{
    return !!this.adService.selectedAd;
  }

  ngOnInit(): void {
    if (this.adService.selectedAd){
      this.adForm.reset(this.adService.selectedAd);
      this.images = this.adService.selectedAd.Images;
    }
  }

}
