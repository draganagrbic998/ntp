import { Component, Input, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { DIALOG_OPTIONS } from 'src/app/constants/dialog';
import { Advertisement } from 'src/app/models/ad';
import { Comment } from 'src/app/models/comment';
import { AdService } from 'src/app/services/ad/ad.service';
import { AuthService } from 'src/app/services/auth/auth.service';
import { CommentService } from 'src/app/services/comment/comment.service';
import { EventService } from 'src/app/services/event/event.service';
import { environment } from 'src/environments/environment';
import { DeleteConfirmationComponent } from '../../shared/controls/delete-confirmation/delete-confirmation.component';

@Component({
  selector: 'app-ad-details',
  templateUrl: './ad-details.component.html',
  styleUrls: ['./ad-details.component.scss']
})
export class AdDetailsComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private adService: AdService,
    private commentService: CommentService,
    private router: Router,
    private dialog: MatDialog,
    private eventService: EventService
  ) { }

  @Input() ad: Advertisement;
  reply: FormControl = new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]);
  replyPending = false;

  edit(): void{
    this.adService.selectedAd = this.ad;
    this.router.navigate([environment.adFormRoute]);
  }

  delete(): void{
    const options: MatDialogConfig = {...DIALOG_OPTIONS, ...{data: () => this.adService.delete(this.ad.ID)}};
    this.dialog.open(DeleteConfirmationComponent, options).afterClosed().subscribe(result => {
      if (result){
        this.adService.announceRefreshData();
      }
    });
  }

  create(): void{
    this.eventService.selectedEvent = null;
    this.router.navigate([`${environment.eventFormRoute}/${this.ad.ID}`]);
  }

  goPage(): void{
    this.router.navigate([`${environment.adPageRoute}/${this.ad.ID}`]);
  }

  get onPage(): boolean{
    return this.router.url.includes('ad-page');
  }

  get id(): number{
    return this.authService.getUser().id;
  }

  sendReply(): void{
    this.replyPending = true;
    this.commentService.save({text: this.reply.value, product_id: this.ad.ID} as Comment).subscribe(() => {
        this.replyPending = false;
        this.reply.reset();
        this.commentService.announceRefreshData();
    });
  }

  ngOnInit(): void {
  }

}
