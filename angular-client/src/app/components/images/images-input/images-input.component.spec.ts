import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ImagesInputComponent } from './images-input.component';

describe('ImagesInputComponent', () => {
  let component: ImagesInputComponent;
  let fixture: ComponentFixture<ImagesInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ImagesInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ImagesInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
