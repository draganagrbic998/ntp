import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { environment } from 'src/environments/environment';
import { AdFormComponent } from './components/ads/ad-form/ad-form.component';
import { AdListComponent } from './components/ads/ad-list/ad-list.component';
import { AdPageComponent } from './components/ads/ad-page/ad-page.component';
import { EventFormComponent } from './components/events/event-form/event-form.component';
import { LoginFormComponent } from './components/user/login-form/login-form.component';
import { ProfileFormComponent } from './components/user/profile-form/profile-form.component';
import { RegisterFormComponent } from './components/user/register-form/register-form.component';
import { AuthGuard } from './guard/auth.guard';

const routes: Routes = [
  {
    path: environment.loginRoute,
    component: LoginFormComponent
  },
  {
    path: environment.registerRoute,
    component: RegisterFormComponent
  },
  {
    path: environment.profileRoute,
    component: ProfileFormComponent,
    canActivate: [AuthGuard]
  },
  {
    path: environment.adListRoute,
    component: AdListComponent,
    canActivate: [AuthGuard]
  },
  {
    path: environment.adFormRoute,
    component: AdFormComponent,
    canActivate: [AuthGuard]
  },
  {
    path: `${environment.adPageRoute}/:id`,
    component: AdPageComponent,
    canActivate: [AuthGuard]
  },
  {
    path: `${environment.eventFormRoute}/:productId`,
    component: EventFormComponent,
    canActivate: [AuthGuard]
  },
  {
    path: '**',
    pathMatch: 'full',
    redirectTo: 'login'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
