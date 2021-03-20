import { Component, Input, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatExpansionPanel } from '@angular/material/expansion';
import { DIALOG_OPTIONS } from 'src/app/constants/dialog';
import { Comment } from 'src/app/models/comment';
import { AuthService } from 'src/app/services/auth/auth.service';
import { CommentService } from 'src/app/services/comment/comment.service';
import { DeleteConfirmationComponent } from '../../shared/controls/delete-confirmation/delete-confirmation.component';

@Component({
  selector: 'app-comment-details',
  templateUrl: './comment-details.component.html',
  styleUrls: ['./comment-details.component.scss']
})
export class CommentDetailsComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private commentService: CommentService,
    private dialog: MatDialog
  ) { }

  @Input() comment: Comment;
  reply: FormControl = new FormControl('', [Validators.required, Validators.pattern(new RegExp('\\S'))]);
  replyPending = false;
  likePending = false;
  dislikePending = false;
  replies: Comment[];

  get id(): number{
    return this.authService.getUser()?.id;
  }

  delete(): void{
    const options: MatDialogConfig = {...DIALOG_OPTIONS, ...{data: () => this.commentService.delete(this.comment.id)}};
    // tslint:disable-next-line: deprecation
    this.dialog.open(DeleteConfirmationComponent, options).afterClosed().subscribe(result => {
      if (result){
        this.commentService.announceRefreshData();
      }
    });
  }

  like(dislike: boolean): void{
    if (!dislike){
      this.likePending = true;
    }
    else{
      this.dislikePending = true;
    }
    // tslint:disable-next-line: deprecation
    this.commentService.like(this.comment.id, dislike).subscribe((response: boolean) => {
      if (!dislike){
        this.likePending = false;
      }
      else{
        this.dislikePending = false;
      }
      if (response){
        if (!dislike){
          this.comment.liked = !this.comment.liked;
          if (this.comment.liked){
            this.comment.likes += 1;
            if (this.comment.disliked){
              this.comment.dislikes -= 1;
              this.comment.disliked = false;
            }
          }
          else{
            this.comment.likes -= 1;
          }
        }
        else{
          this.comment.disliked = !this.comment.disliked;
          if (this.comment.disliked){
            this.comment.dislikes += 1;
            if (this.comment.liked){
              this.comment.likes -= 1;
              this.comment.liked = false;
            }
          }
          else{
            this.comment.dislikes -= 1;
          }
        }
      }
    });
  }

  sendReply(replies: MatExpansionPanel): void{
    this.replyPending = true;
    this.commentService.save({text: this.reply.value, parent_id: this.comment.id,
      // tslint:disable-next-line: deprecation
      product_id: this.comment.product_id} as Comment).subscribe(() => {
        this.replyPending = false;
        this.reply.reset();
        replies.open();
    });
  }

  fetchReplies(): void{
    // tslint:disable-next-line: deprecation
    this.commentService.replies(this.comment.id).subscribe((comments: Comment[]) => {
      this.replies = comments;
    });
  }

  ngOnInit(): void {
  }

}
