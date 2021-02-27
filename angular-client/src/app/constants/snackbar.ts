import { MatSnackBarConfig } from '@angular/material/snack-bar';

export const SNACKBAR_CLOSE = 'Close';
export const SNACKBAR_ERROR = 'An error occured! Try again.';
export const DELETE_SUCCESS = 'Item successfully removed!';
export const DELETE_ERROR = 'Item cannot be removed!';

export const SNACKBAR_SUCCESS_OPTIONS: MatSnackBarConfig = {
    horizontalPosition: 'center',
    verticalPosition: 'top',
    panelClass: 'snackbar-success',
    duration: 2000
};

export const SNACKBAR_ERROR_OPTIONS: MatSnackBarConfig = {
    horizontalPosition: 'center',
    verticalPosition: 'top',
    panelClass: 'snackbar-error',
    duration: 2000
};
