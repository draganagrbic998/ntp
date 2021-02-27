import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ProfileComponent } from 'src/app/components/user/profile/profile.component';
import { AdService } from 'src/app/services/ad/ad.service';
import { AuthService } from 'src/app/services/auth/auth.service';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {

  constructor(
    private authService: AuthService,
    public adService: AdService,
    private router: Router,
    private dialog: MatDialog
  ) { }

  search: FormControl = new FormControl('');

  create(): void{
    this.adService.selectedAd = null;
    this.router.navigate([environment.adFormRoute]);
  }

  get route(): string{
    return this.router.url.substr(1);
  }

  openProfile(): void{
    const options: MatDialogConfig = {
      panelClass: 'no-padding',
      backdropClass: 'cdk-overlay-transparent-backdrop',
      position: {
          top: '50px',
          right: '30px'
      }
    };
    this.dialog.open(ProfileComponent, options);
  }

  signOut(): void{
    this.authService.deleteUser();
    this.router.navigate([environment.loginRoute]);
  }

  ngOnInit(): void {
  }

}
