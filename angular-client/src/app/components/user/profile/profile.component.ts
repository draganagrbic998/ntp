import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth/auth.service';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private router: Router,
    private dialogRef: MatDialogRef<ProfileComponent>
  ) { }

  profile: User = this.authService.getUser();

  edit(): void{
    this.dialogRef.close();
    this.router.navigate([environment.profileRoute]);
  }

  ngOnInit(): void {
  }

}
