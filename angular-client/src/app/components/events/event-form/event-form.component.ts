import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute } from '@angular/router';
import { SNACKBAR_CLOSE, SNACKBAR_ERROR, SNACKBAR_ERROR_OPTIONS, SNACKBAR_SUCCESS_OPTIONS } from 'src/app/constants/snackbar';
import { Event } from 'src/app/models/event';
import { Image } from 'src/app/models/image';
import { EventService } from 'src/app/services/event/event.service';
import { DateValidatorService } from './date-validator.service';

@Component({
  selector: 'app-event-form',
  templateUrl: './event-form.component.html',
  styleUrls: ['./event-form.component.scss']
})
export class EventFormComponent implements OnInit {

  constructor(
    private eventService: EventService,
    private snackBar: MatSnackBar,
    private route: ActivatedRoute,
    public location: Location,
    private dateValidator: DateValidatorService
  ) { }

  savePending = false;
  eventForm: FormGroup = new FormGroup({
    Category: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    Name: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    From: new FormControl('', [Validators.required, this.dateValidator.dateValidator()]),
    To: new FormControl('', [Validators.required, this.dateValidator.dateValidator()]),
    Place: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]),
    Description: new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))])
  });
  images: Image[] = [];

  save(): void{
    if (this.eventForm.invalid){
      return;
    }
    this.savePending = true;
    this.eventService.save({...this.eventForm.value, ...{Images: this.images,
      ID: this.eventService.selectedEvent?.ID, ProductID: +this.route.snapshot.params.productId}}).subscribe(
      (event: Event) => {
        this.savePending = false;
        if (event){
          this.snackBar.open('Event saved!', SNACKBAR_CLOSE, SNACKBAR_SUCCESS_OPTIONS);
          this.eventService.announceRefreshData();
        }
        else{
          this.snackBar.open(SNACKBAR_ERROR, SNACKBAR_CLOSE, SNACKBAR_ERROR_OPTIONS);
        }
      }
    );
  }

  get edit(): boolean{
    return !!this.eventService.selectedEvent;
  }

  ngOnInit(): void {
    if (this.eventService.selectedEvent){
      this.eventForm.reset(this.eventService.selectedEvent);
      this.images = this.eventService.selectedEvent.Images;
    }
  }

}
