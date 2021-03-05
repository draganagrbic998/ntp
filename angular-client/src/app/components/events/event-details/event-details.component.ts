import { Component, Input, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { DIALOG_OPTIONS } from 'src/app/constants/dialog';
import { Event } from 'src/app/models/event';
import { AuthService } from 'src/app/services/auth/auth.service';
import { EventService } from 'src/app/services/event/event.service';
import { environment } from 'src/environments/environment';
import { DeleteConfirmationComponent } from '../../shared/controls/delete-confirmation/delete-confirmation.component';

@Component({
  selector: 'app-event-details',
  templateUrl: './event-details.component.html',
  styleUrls: ['./event-details.component.scss']
})
export class EventDetailsComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private eventService: EventService,
    private router: Router,
    private dialog: MatDialog
  ) { }

  @Input() event: Event;

  edit(): void{
    this.eventService.selectedEvent = this.event;
    this.router.navigate([`${environment.eventFormRoute}/${this.event.ProductID}`]);
  }

  delete(): void{
    const options: MatDialogConfig = {...DIALOG_OPTIONS, ...{data: () => this.eventService.delete(this.event.ID)}};
    this.dialog.open(DeleteConfirmationComponent, options).afterClosed().subscribe(result => {
      if (result){
        this.eventService.announceRefreshData();
      }
    });
  }

  get id(): number{
    return this.authService.getUser().id;
  }

  ngOnInit(): void {
  }

}
