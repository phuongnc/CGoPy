
//
//  filter.h
//  go2py
//
//  Created by PhuongNC on 4/30/20.
//  Copyright © 2020 PhuongNC. All rights reserved.
//

#ifndef PY_FILTER_H
#define PY_FILTER_H

#ifdef __cplusplus
extern "C" {
#endif

extern void Py_Load();
extern void Py_Unload();
extern char *Py_Filter_Text(const char *text);

#ifdef __cplusplus
}
#endif

#endif // PY_FILTER_H
