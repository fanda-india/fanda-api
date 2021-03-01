-- seed data

-- name: insert-admin-user
INSERT OR REPLACE INTO users (id, login_name, email, mobile_number, first_name, last_name, user_password, is_enabled, created_at, is_reset_pwd)
VALUES (1, 'admin', 'admin@fanda2.com', '9999999999', 'System', 'Administrator', '123', true, '2020-01-01', false);

-- name: insert-ledger-groups
BEGIN;
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (100, 'AUSPIC', 'AUSPICIOUS ACCOUNTS', 1, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (200, 'CAPITAL', 'CAPITAL ACCOUNTS', 1, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (300, 'CURRLIA', 'CURRENT LIABILITIES', 2, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (400, 'LOANS', 'LOANS (LIABILITY)', 2, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (500, 'FIXAS', 'FIXED ASSETS', 1, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (600, 'INVEST', 'INVESTMENTS (ASSET)', 1, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (700, 'CURRAS', 'CURRENT ASSETS', 1, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (800, 'REVACC', 'REVENUE ACCOUNTS', 3, NULL);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (310, 'TAXES', 'Duties and Taxes', 2, 300);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (320, 'SUNCRS', 'Sundry Creditors', 2, 300);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (710, 'BANKS', 'Bank Accounts', 1, 700);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (720, 'CASH', 'Cash-in-hand', 1, 700);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (730, 'DEPOSIT', 'Deposits', 1, 700);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (740, 'ADVANCE', 'Advances', 1, 700);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (750, 'STOCK', 'Stock-in-hand', 1, 700);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (760, 'SUNDRS', 'Sundry Debtors', 1, 700);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (810, 'PURCHASE', 'Purchase Accounts', 3, 800);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (820, 'SALES', 'Sales Accounts', 3, 800);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (830, 'INCOME', 'Income Accounts', 4, 800);
    INSERT OR REPLACE INTO ledger_groups (id, code, group_name, group_type, parent_id)
    VALUES (840, 'EXPENSE', 'Expenses Accounts', 5, 800);
END;

/*('00000000-0000-0000-0000-000000000001', 'MISCEXP', 'MISCELLANEOUS EXPENSES (ASSET)', NULL, 1, NULL),
('00000000-0000-0000-0000-000000000001', 'BR/DIV', 'BRANCHES/DIVISIONS', NULL, 1, NULL),
('00000000-0000-0000-0000-000000000001', 'PRO/LOS', 'PROFIT & LOSS', NULL, 1, NULL),
('00000000-0000-0000-0000-000000000001', 'SUSPACC', 'SUSPENSE ACCOUNTS', NULL, 1, NULL),*/