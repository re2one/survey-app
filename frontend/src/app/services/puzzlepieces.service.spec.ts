import { TestBed } from '@angular/core/testing';

import { PuzzlepiecesService } from './puzzlepieces.service';

describe('PuzzlepiecesService', () => {
  let service: PuzzlepiecesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PuzzlepiecesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
