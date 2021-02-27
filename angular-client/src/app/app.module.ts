import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { LoginFormComponent } from './components/user/login-form/login-form.component';
import { RegisterFormComponent } from './components/user/register-form/register-form.component';
import { ProfileFormComponent } from './components/user/profile-form/profile-form.component';

import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { AuthInterceptor } from './interceptor/auth.interceptor';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatSelectModule } from '@angular/material/select';
import { ImagesInputComponent } from './components/images/images-input/images-input.component';
import { CarouselComponent } from './components/images/carousel/carousel.component';
import { ImageInputComponent } from './components/images/image-input/image-input.component';
import { CarouselModule } from 'ng-uikit-pro-standard';
import { MatDialogModule } from '@angular/material/dialog';
import { EventDetailsComponent } from './components/events/event-details/event-details.component';
import { EventListComponent } from './components/events/event-list/event-list.component';
import { EventFormComponent } from './components/events/event-form/event-form.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { CommentListComponent } from './components/comments/comment-list/comment-list.component';
import { CommentDetailsComponent } from './components/comments/comment-details/comment-details.component';
import { MatExpansionModule } from '@angular/material/expansion';
import { FormContainerComponent } from './components/shared/containers/form-container/form-container.component';
import { SpacerContainerComponent } from './components/shared/containers/spacer-container/spacer-container.component';
import { CenterContainerComponent } from './components/shared/containers/center-container/center-container.component';
import { PreloaderComponent } from './components/shared/loaders/preloader/preloader.component';
import { SpinnerButtonComponent } from './components/shared/loaders/spinner-button/spinner-button.component';
import { AdDetailsComponent } from './components/ads/ad-details/ad-details.component';
import { AdListComponent } from './components/ads/ad-list/ad-list.component';
import { AdFormComponent } from './components/ads/ad-form/ad-form.component';
import { BoldTextComponent } from './components/shared/containers/bold-text/bold-text.component';
import { ToolbarComponent } from './components/shared/controls/toolbar/toolbar.component';
import { PaginatorComponent } from './components/shared/controls/paginator/paginator.component';
import { DeleteConfirmationComponent } from './components/shared/controls/delete-confirmation/delete-confirmation.component';
import { AdPageComponent } from './components/ads/ad-page/ad-page.component';
import { MatMenuModule } from '@angular/material/menu';
import { ProfileComponent } from './components/user/profile/profile.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginFormComponent,
    RegisterFormComponent,
    ProfileFormComponent,
    FormContainerComponent,
    SpacerContainerComponent,
    CenterContainerComponent,
    PreloaderComponent,
    SpinnerButtonComponent,
    AdDetailsComponent,
    AdListComponent,
    AdFormComponent,
    BoldTextComponent,
    ToolbarComponent,
    PaginatorComponent,
    ImagesInputComponent,
    CarouselComponent,
    ImageInputComponent,
    DeleteConfirmationComponent,
    AdPageComponent,
    EventDetailsComponent,
    EventListComponent,
    EventFormComponent,
    CommentListComponent,
    CommentDetailsComponent,
    ProfileComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    NgbModule,

    MatCardModule,
    MatProgressSpinnerModule,
    MatSnackBarModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    MatTooltipModule,
    MatSelectModule,
    CarouselModule,
    MatDialogModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatExpansionModule,
    MatMenuModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
