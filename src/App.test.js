import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders the app', () => {
  const div = document.createElement('div');
  render(<App />, div);
});
