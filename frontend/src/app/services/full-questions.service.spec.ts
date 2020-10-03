import { TestBed } from '@angular/core/testing';

import { FullQuestionsService } from './full-questions.service';

describe('FullQuestionsService', () => {
  let service: FullQuestionsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FullQuestionsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
