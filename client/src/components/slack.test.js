import { SlackConnect } from './SlackConnect';
import { render, waitFor } from '@testing-library/react';
import axios from 'axios'

let el;
beforeEach(() => {
    el = render(<SlackConnect />)
});

jest.mock('axios');
axios.get.mockReturn1qValue({ data: true });

describe("on render", () => {
    afterEach(() => {
        jest.clearAllMocks();
    })
    it('should call axios', async () => {
        await waitFor(() => expect(axios.get).toHaveBeenCalled());
    });
    it('should have a state of "not connected"', () => {
        const { queryByText } = el;
        const label = queryByText('Not Connected: Slack');
        expect(label).not.toBe(null);
    });
    it('should update the label on connecting to slack', async () => {
        const { queryByText } = el;
        const label = queryByText('Connected: Slack');
        await waitFor(() => expect(label).not.toBe(null));
    })
})
