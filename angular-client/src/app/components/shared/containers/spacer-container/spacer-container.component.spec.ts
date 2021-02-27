import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SpacerContainerComponent } from './spacer-container.component';

describe('SpacerContainerComponent', () => {
  let component: SpacerContainerComponent;
  let fixture: ComponentFixture<SpacerContainerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SpacerContainerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SpacerContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
