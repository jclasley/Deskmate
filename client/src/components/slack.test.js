import { SlackConnect } from './SlackConnect';
import { render, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event'
import axios from 'axios'

let el;
jest.mock('axios');


describe("on render", () => {
    beforeEach(() => {
        axios.get.mockResolvedValue({ data: true });
    });
    it('should call axios', async () => {
        el = render(<SlackConnect />)
        await waitFor(() => expect(axios.get).toHaveBeenCalled());
    });
    it('should have a state of "not connected"', () => {
        axios.get.mockResolvedValue({ data: false });
        el = render(<SlackConnect />)
        const { queryByText } = el;
        const label = queryByText('Not Connected: Slack');
        expect(label).not.toBe(null);
    });
    it('should update the label on connecting to slack', async () => {
        el = render(<SlackConnect />)
        const { queryByText } = el;
        const label = queryByText('Connected: Slack');
        expect(axios.get).toHaveBeenCalled();
        await waitFor(() => expect(label).not.toBe(null));
    })
});

describe('on click', () => {
    beforeEach(() => {
        axios.get.mockResolvedValue({ data: false });
        el = render(<SlackConnect />)
    })
    it('should call axios again on click', async () => {
        const { getByRole } = el;
        const b = getByRole('button');
        userEvent.click(b);
        // should be called once on init and once on click
        await waitFor(() => expect(axios.get).toHaveBeenCalledTimes(2));
    });
})