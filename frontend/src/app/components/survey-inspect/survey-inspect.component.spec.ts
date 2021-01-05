import {async, ComponentFixture, TestBed} from '@angular/core/testing';

import {SurveyInspectComponent} from './survey-inspect.component';

describe('SurveyInspectComponent', () => {
  let component: SurveyInspectComponent;
  let fixture: ComponentFixture<SurveyInspectComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [SurveyInspectComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SurveyInspectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
