import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Observable } from 'rxjs';
import { DELETE_ERROR, DELETE_SUCCESS, SNACKBAR_CLOSE, SNACKBAR_ERROR_OPTIONS, SNACKBAR_SUCCESS_OPTIONS } from 'src/app/constants/snackbar';

@Component({
  selector: 'app-delete-confirmation',
  templateUrl: './delete-confirmation.component.html',
  styleUrls: ['./delete-confirmation.component.scss']
})
export class DeleteConfirmationComponent implements OnInit {

  constructor(
    @Inject(MAT_DIALOG_DATA) private deleteFunction: () => Observable<boolean>,
    public dialogRef: MatDialogRef<DeleteConfirmationComponent>,
    private snackBar: MatSnackBar
  ) { }

  deletePending = false;

  confirm(): void{
    this.deletePending = true;
    this.deleteFunction().subscribe(
      (param: boolean) => {
        this.deletePending = false;
        if (param){
          this.snackBar.open(DELETE_SUCCESS, SNACKBAR_CLOSE, SNACKBAR_SUCCESS_OPTIONS);
          this.dialogRef.close(param);
        }
        else{
          this.snackBar.open(DELETE_ERROR, SNACKBAR_CLOSE, SNACKBAR_ERROR_OPTIONS);
        }
      }
    );
  }

  ngOnInit(): void {
  }

}
