package page_generator

// // First step
// func (m *mainView) handleEnter() (*mainView, tea.Cmd) {
// 	if m.start {
// 		return m, nil
// 	}

// 	m.start = true
// 	// start first step
// 	// TaskFlow first task is -1, so to start first task, we need to use [nextTask]
// 	cmd := m.taskFlow.NextTask()
// 	m.taskFlow.SetCurrentStatus(component.InProgress)

// 	return m, tea.Batch(m.logic.GetFutures(), cmd)
// }

// // When first step end. It also init second step
// func (m *mainView) handleFuturesMsg(msg FuturesMsg) (tea.Model, tea.Cmd) {
// 	if msg.err != nil {
// 		// m.steps[0].err = msg.err
// 		// return m, nil
// 		// do not start next task if error
// 	}

// 	var (
// 		cmd  tea.Cmd
// 		cmds []tea.Cmd
// 	)

// 	m.futures = msg.futures

// 	// start next task (dexscreener)
// 	cmd = m.logic.GetPairs(m.futures)
// 	cmds = append(cmds, cmd)

// 	// next task vizually
// 	m.taskFlow.SetCurrentStatus(component.Done)
// 	cmd = m.taskFlow.NextTask()
// 	m.taskFlow.SetCurrentStatus(component.InProgress)
// 	cmds = append(cmds, cmd)

// 	return m, tea.Batch(cmds...)
// }

// func (m *mainView) handleDsMsg(msg DexscreenerMsg) (tea.Model, tea.Cmd) {
// 	if msg.err != nil {
// 		// do smth
// 	}

// 	m.pairs = msg.pairs

// 	// move to the next task
// 	m.taskFlow.SetCurrentStatus(component.Done)
// 	cmd := m.taskFlow.NextTask()
// 	m.taskFlow.SetCurrentStatus(component.InProgress)

// 	return m, cmd
// }

// // save pairs / tokens to the database
// func (m *mainView) handleTasksEnd() (tea.Model, tea.Cmd) {
// 	m.taskFlow.SetCurrentStatus(component.Done)

// 	// to understand next logic you need to know next:
// 	//  - first step of generation is to fetch all existing tokens (on futures) from mexc
// 	//  - second step is to fetch pairs according to tokens from the 1 step
// 	//  - but sometimes there are no pairs for some tokens (too low volume for example)
// 	//    so for this situation we also need to filter tokens by pairs, to not save bad tokens, that didnot found pair
// 	ctx := context.Background()
// 	for _, pair := range m.pairs {
// 		// find token. Quote tokens we cannot find, due to the mexc api spec. (Mexc just does not provide contract for L1 tokens)
// 		var base market.Token
// 		for _, token := range m.futures {
// 			if pair.BaseToken.Address == token.Address {
// 				base = token
// 			}
// 		}

// 		// todo: quote token will not be found always
// 		// save or find exist
// 		baseId, err := m.tokenRepo.FindOrCreate(ctx, base)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, nil
// 		}

// 		quoteId, err := m.tokenRepo.FindOrCreate(ctx, pair.QuoteToken)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, nil
// 		}

// 		pair.BaseToken = market.Token{}
// 		pair.QuoteToken = market.Token{}
// 		_, err = m.pairRepo.FindOrCreate(ctx, pair, baseId, quoteId)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, nil
// 		}
// 	}

// 	return m, nil
// }
