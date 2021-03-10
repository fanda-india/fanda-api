INSERT INTO users (user_name, email, mobile_number, first_name, last_name, password, active, created_at, is_password_reset)
VALUES
('admin', 'admin@fanda2.com', '9999999999', 'System', 'Administrator', '123', true, '2020-01-01', false);

INSERT INTO ledger_groups (id, code, name, group_type, parent_id) VALUES
(100, 'AUSPIC', 'AUSPICIOUS ACCOUNTS', 1, NULL),
(200, 'CAPITAL', 'CAPITAL ACCOUNTS', 1, NULL),
(300, 'CURRLIA', 'CURRENT LIABILITIES', 2, NULL),
(400, 'LOANS', 'LOANS (LIABILITY)', 2, NULL),
(500, 'FIXAS', 'FIXED ASSETS', 1, NULL),
(600, 'INVEST', 'INVESTMENTS (ASSET)', 1, NULL),
(700, 'CURRAS', 'CURRENT ASSETS', 1, NULL),
(800, 'REVACC', 'REVENUE ACCOUNTS', 3, NULL),

(310, 'TAXES', 'Duties and Taxes', 2, 300),
(320, 'SUNCRS', 'Sundry Creditors', 2, 300),

(710, 'BANKS', 'Bank Accounts', 1, 700),
(720, 'CASH', 'Cash-in-hand', 1, 700),
(730, 'DEPOSIT', 'Deposits', 1, 700),
(740, 'ADVANCE', 'Advances', 1, 700),
(750, 'STOCK', 'Stock-in-hand', 1, 700),
(760, 'SUNDRS', 'Sundry Debtors', 1, 700),

(810, 'PURCHASE', 'Purchase Accounts', 3, 800),
(820, 'SALES', 'Sales Accounts', 3, 800),
(830, 'INCOME', 'Income Accounts', 4, 800),
(840, 'EXPENSE', 'Expenses Accounts', 5, 800);
