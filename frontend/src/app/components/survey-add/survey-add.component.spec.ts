import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SurveyAddComponent } from './survey-add.component';

describe('SurveyAddComponent', () => {
  let component: SurveyAddComponent;
  let fixture: ComponentFixture<SurveyAddComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SurveyAddComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SurveyAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
